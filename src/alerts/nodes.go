package alerts

import (
	"encoding/json"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"log"
	"observer/src/models"
)

func NodesCheck() {
	log.Println("[Alerts] - NodesCheck")

	nodeList, err := resources.RetrieveNodesList()
	if err != nil {
		log.Println("[Alerts] - Error on resources.RetrieveNodesList")
		return
	}

	for _, nodeInList := range nodeList.Items{
		if isNodeUnhealthy(nodeInList) {
			nodeInfo := &models.Node{
				Name:           nodeInList.Name,
				MemoryPressure: string(nodeInList.Status.Conditions[0].Status),
				DiskPressure:   string(nodeInList.Status.Conditions[1].Status),
				PIDPressure:    string(nodeInList.Status.Conditions[2].Status),
				Ready:          string(nodeInList.Status.Conditions[3].Status),
			}
			info, _ := json.Marshal(nodeInfo)
			fmt.Println(string(info))
		}
	}
}

func isNodeUnhealthy(node v1.Node) bool {
	return string(node.Status.Conditions[0].Status) == "True" || string(node.Status.Conditions[1].Status) == "True" ||
		string(node.Status.Conditions[2].Status) == "True" ||
		string(node.Status.Conditions[3].Status) == "False"
}
