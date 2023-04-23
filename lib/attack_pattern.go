package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"thiapp/neptune_loader/ent"
	"time"
)

type BundleData struct {
	Type        string                   `json:"type"`
	ID          string                   `json:"id"`
	Objects     []map[string]interface{} `json:"objects"`
	SpecVersion string                   `json:"spec_version"`
}

func GetBundleFromUrl(downloadSrcUrl string) (*BundleData, error) {
	response, err := http.Get(downloadSrcUrl)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logrus.Fatalln(err)
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	bundle := &BundleData{}
	err = json.Unmarshal(body, &bundle)

	return bundle, nil
}

func PushAttackPatternsToDB(ctx context.Context, db *ent.Client, objects []map[string]interface{}, stripDollar bool) {
	for _, row := range objects {
		if row["type"] == "attack-pattern" {
			item := create(db, row, stripDollar)
			_, err := item.Save(ctx)

			if err != nil {
				logrus.Fatalln(err)
			}
		}
	}
}

func create(db *ent.Client, row map[string]interface{}, stripDollar bool) *ent.AttackPatternCreate {
	return db.AttackPattern.Create().
		SetID(getString(row, "id")).
		SetType(getString(row, "type")).
		SetName(getString(row, "name")).
		SetNillableAliases(getStringOrNil(row, "aliases", stripDollar)).
		SetNillableKillChainPhases(getStringOrNil(row, "kill_chain_phases", stripDollar)).
		SetNillableConfidence(getStringOrNil(row, "confidence", stripDollar)).
		SetNillableLabels(getStringOrNil(row, "labels", stripDollar)).
		SetNillableLang(getStringOrNil(row, "lang", stripDollar)).
		SetNillableObjectMarkingRefs(getStringOrNil(row, "object_marking_refs", stripDollar)).
		SetNillableGranularMarkings(getStringOrNil(row, "granular_markings", stripDollar)).
		SetNillableExtensions(getStringOrNil(row, "extensions", stripDollar)).
		SetNillableExternalReferences(getStringOrNil(row, "external_references", stripDollar)).
		SetNillableSpecVersion(getStringOrNil(row, "spec_version", stripDollar)).
		SetNillableCreatedByRef(getStringOrNil(row, "created_by_ref", stripDollar)).
		SetNillableRevoked(getBoolOrNil(row, "revoked")).
		SetNillableCreated(getTimeOrNil(row, "created")).
		SetNillableModified(getTimeOrNil(row, "modified")).
		SetNillableDescription(getStringOrNil(row, "description", stripDollar)).
		SetNillableXMitreDataSources(getStringOrNil(row, "x_mitre_data_sources", stripDollar)).
		SetNillableXMitreAttackSpecVersion(getStringOrNil(row, "x_mitre_attack_spec_version", stripDollar)).
		SetNillableXMitreDeprecated(getBoolOrNil(row, "x_mitre_deprecated")).
		SetNillableXMitreDetection(getStringOrNil(row, "x_mitre_detection", stripDollar)).
		SetNillableXMitreDefenseBypassed(getStringOrNil(row, "x_mitre_defense_bypassed", stripDollar)).
		SetNillableXMitreDomains(getStringOrNil(row, "x_mitre_domains", stripDollar)).
		SetNillableXMitreIsSubtechnique(getBoolOrNil(row, "x_mitre_is_subtechnique")).
		SetNillableXMitreModifiedByRef(getStringOrNil(row, "x_mitre_modified_by_ref", stripDollar)).
		SetNillableXMitrePermissionsRequired(getStringOrNil(row, "x_mitre_permissions_required", stripDollar)).
		SetNillableXMitrePlatforms(getStringOrNil(row, "x_mitre_platforms", stripDollar)).
		SetNillableXMitreRemoteSupport(getBoolOrNil(row, "x_mitre_remote_support")).
		SetNillableXMitreSystemRequirements(getStringOrNil(row, "x_mitre_system_requirements", stripDollar)).
		SetNillableXMitreVersion(getStringOrNil(row, "x_mitre_version", stripDollar)).
		SetNillableXMitreContributors(getStringOrNil(row, "x_mitre_contributors", stripDollar))
}

func getString(row map[string]interface{}, key string) string {
	iValue := row[key]
	if iValue == nil {
		return ""
	}
	return row[key].(string)
}

func getStringOrNil(row map[string]interface{}, key string, stripDollar bool) *string {
	iValue := row[key]
	valueType := fmt.Sprintf("%v", reflect.ValueOf(iValue).Kind())
	if iValue == nil {
		return nil
	}
	var value string
	if valueType == "slice" {
		value = fmt.Sprintf("%v", iValue)
	} else {
		value = iValue.(string)
	}
	if stripDollar {
		value = strings.ReplaceAll(value, "$", "")
	}

	return &value
}

func getBoolOrNil(row map[string]interface{}, key string) *bool {
	iValue := row[key]
	if iValue == nil {
		return nil
	}
	value := iValue.(bool)
	return &value
}

func getTimeOrNil(row map[string]interface{}, key string) *time.Time {
	iValue := row[key]
	if iValue == nil {
		return nil
	}
	valueStr := fmt.Sprintf("%v", iValue)
	timeValue, _ := time.Parse(time.RFC3339, valueStr)
	return &timeValue
}
