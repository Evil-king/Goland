package main

import (
	"context"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
)

var (
	versionId string = "kkkkk"
)

// Alertmanager YAML Implement
type AlertYAMLImpl struct {
	Client     *grt.Client
	ServerAddr string
	VersionId  string
}

// Start code here
func CheckRulesFiles(ctx context.Context) error {
	// ServiceName lookup
	nsh, err := net.LookupNS("amtrans")
	if err != nil {
		return err
	}
	nscp := cap(nsh)
	ag := AlertYAMLImpl{
		Client:     common.Nccl(),
		ServerAddr: fmt.Sprintf("http://%s:8080/sub/alerts_rules", nsh[rand.Intn(nscp)].Host),
		//ServerAddr: fmt.Sprintf("http://%s:8080/sub/alerts_rules", "localhost"),
	}
	// Send currently version to WebServer to check
	if !ag.verifyVersionId(ctx, versionId) {
		return errors.New("Rules version id is same, no need to write down")
	}
	// Got lastest config file yaml write down and reload
	if err := ag.writeDown(ctx); err != nil {
		return err
	}
	// Reload alertmanager by API
	rep, err := common.AlmReload(ctx)
	if err != nil {
		return err
	}
	return nil
}

func writeRulesYaml(data []byte) (int, error) {
	rulesPath := viper.GetString("AlertmanagerRulesPath")
	if rulesPath == "" {
		rulesPath = "/etc/alertmanager/rules"
	}
	//rulesPath := "./rules.yaml"
	f, err := os.OpenFile(rulesPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return 0, err
	}
	return f.Write(data)
}

func (al AlertYAMLImpl) verifyVersionId(ctx context.Context, versionId string) bool {
	var rp2 common.HTTPRp
	rq, err := grt.NewRequestWithContext(ctx, http.MethodGet, al.ServerAddr, nil)
	if err != nil {
		return false
	}
	rp, err := al.Client.Do(rq)
	if err != nil {
		return false
	}
	if rp.StatusCode != http.StatusOK {
		return false
	}
	rbo, err := io.ReadAll(rp.Body)
	if err != nil {
		return false
	}
	if err := jsoniter.Unmarshal(rbo, &rp2); err != nil {
		return false
	}
	log.Debug(rp2)
	if rp2.Data != versionId {
		al.VersionId = versionId
		log.Info("version id is different, than replace")
		return true
	}
	log.Warn("version id is same, no need to replace")
	return false
}

func (al AlertYAMLImpl) writeDown(ctx context.Context) error {
	var w2 common.HTTPRp
	var w3 common.DSKeyrequest
	pbody, err := jsoniter.Marshal(&common.PostVersionId{
		VersionId: versionId,
	})
	if err != nil {
		return err
	}
	rq, err := grt.NewRequestWithContext(ctx, http.MethodPost, al.ServerAddr, pbody)
	if err != nil {
		return err
	}
	rp, err := common.Nccl().Do(rq)
	if err != nil {
		return err
	}
	rw2, err := io.ReadAll(rp.Body)
	if err != nil {
		return err
	}
	if err := jsoniter.Unmarshal(rw2, &w2); err != nil {
		return err
	}
	if err := jsoniter.UnmarshalFromString(w2.Data, &w3); err != nil {
		return err
	}
	data, err := buildRulesYAML(&w3)
	if err != nil {
		return err
	}
	log.Info("write rules yaml")
	_, err := writeRulesYaml(data)
	if err != nil {
		return err
	}
	return nil
}

// Rebuild rules yaml from keyRequest struct
func buildRulesYAML(ain *common.DSKeyrequest) ([]byte, error) {
	var (
		appg []common.WGroups
		g2   common.WGroups
		r2   common.WRules
		l2   common.WLabels
		a2   common.WAnnotations
	)
	l2.Severity = "page"
	for _, v := range ain.KeyRequest {
		for i2 := range v.MetricsQl {
			if v.MetricsQl[i2].Latency != "" {
				a2.Summary = fmt.Sprintf("KeyRequestLatency: %s", strings.ReplaceAll(v.URLPath, "/", " "))
				r2.Alert = fmt.Sprintf("%s_latency", v.MetricsQl[i2].SubName)
				r2.Expr = v.MetricsQl[i2].Latency
				r2.For = "0m"
				r2.Labels = l2
				r2.Annotations = a2
				g2.Rules = append(g2.Rules, r2)
			}
			// ...
			appg = append(appg, g2)
		}
	}
	return yaml.Marshal(common.WAlertRules{
		Groups: appg,
	})
}

func convertdash2space(input string) string {
	return strings.ReplaceAll(input, "/", " ")
}
