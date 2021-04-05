// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// register flags to command
func registerModelSwarmInfoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSwarmInfoCluster(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSwarmInfoControlAvailable(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSwarmInfoError(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSwarmInfoLocalNodeState(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSwarmInfoManagers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSwarmInfoNodeAddr(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSwarmInfoNodeID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSwarmInfoNodes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSwarmInfoRemoteManagers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSwarmInfoCluster(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var clusterFlagName string
	if cmdPrefix == "" {
		clusterFlagName = "Cluster"
	} else {
		clusterFlagName = fmt.Sprintf("%v.Cluster", cmdPrefix)
	}

	registerModelSwarmInfoFlags(depth+1, clusterFlagName, cmd)

	return nil
}

func registerSwarmInfoControlAvailable(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	controlAvailableDescription := ``

	var controlAvailableFlagName string
	if cmdPrefix == "" {
		controlAvailableFlagName = "ControlAvailable"
	} else {
		controlAvailableFlagName = fmt.Sprintf("%v.ControlAvailable", cmdPrefix)
	}

	var controlAvailableFlagDefault bool

	_ = cmd.PersistentFlags().Bool(controlAvailableFlagName, controlAvailableFlagDefault, controlAvailableDescription)

	return nil
}

func registerSwarmInfoError(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	errorDescription := ``

	var errorFlagName string
	if cmdPrefix == "" {
		errorFlagName = "Error"
	} else {
		errorFlagName = fmt.Sprintf("%v.Error", cmdPrefix)
	}

	var errorFlagDefault string

	_ = cmd.PersistentFlags().String(errorFlagName, errorFlagDefault, errorDescription)

	return nil
}

func registerSwarmInfoLocalNodeState(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primitive LocalNodeState LocalNodeState is not supported by go-swagger cli yet

	return nil
}

func registerSwarmInfoManagers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	managersDescription := `Total number of managers in the swarm.`

	var managersFlagName string
	if cmdPrefix == "" {
		managersFlagName = "Managers"
	} else {
		managersFlagName = fmt.Sprintf("%v.Managers", cmdPrefix)
	}

	var managersFlagDefault int64

	_ = cmd.PersistentFlags().Int64(managersFlagName, managersFlagDefault, managersDescription)

	return nil
}

func registerSwarmInfoNodeAddr(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nodeAddrDescription := `IP address at which this node can be reached by other nodes in the
swarm.
`

	var nodeAddrFlagName string
	if cmdPrefix == "" {
		nodeAddrFlagName = "NodeAddr"
	} else {
		nodeAddrFlagName = fmt.Sprintf("%v.NodeAddr", cmdPrefix)
	}

	var nodeAddrFlagDefault string

	_ = cmd.PersistentFlags().String(nodeAddrFlagName, nodeAddrFlagDefault, nodeAddrDescription)

	return nil
}

func registerSwarmInfoNodeID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nodeIdDescription := `Unique identifier of for this node in the swarm.`

	var nodeIdFlagName string
	if cmdPrefix == "" {
		nodeIdFlagName = "NodeID"
	} else {
		nodeIdFlagName = fmt.Sprintf("%v.NodeID", cmdPrefix)
	}

	var nodeIdFlagDefault string

	_ = cmd.PersistentFlags().String(nodeIdFlagName, nodeIdFlagDefault, nodeIdDescription)

	return nil
}

func registerSwarmInfoNodes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nodesDescription := `Total number of nodes in the swarm.`

	var nodesFlagName string
	if cmdPrefix == "" {
		nodesFlagName = "Nodes"
	} else {
		nodesFlagName = fmt.Sprintf("%v.Nodes", cmdPrefix)
	}

	var nodesFlagDefault int64

	_ = cmd.PersistentFlags().Int64(nodesFlagName, nodesFlagDefault, nodesDescription)

	return nil
}

func registerSwarmInfoRemoteManagers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: RemoteManagers []*PeerNode array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSwarmInfoFlags(depth int, m *models.SwarmInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, clusterAdded := retrieveSwarmInfoClusterFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clusterAdded

	err, controlAvailableAdded := retrieveSwarmInfoControlAvailableFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || controlAvailableAdded

	err, errorAdded := retrieveSwarmInfoErrorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorAdded

	err, localNodeStateAdded := retrieveSwarmInfoLocalNodeStateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || localNodeStateAdded

	err, managersAdded := retrieveSwarmInfoManagersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || managersAdded

	err, nodeAddrAdded := retrieveSwarmInfoNodeAddrFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nodeAddrAdded

	err, nodeIdAdded := retrieveSwarmInfoNodeIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nodeIdAdded

	err, nodesAdded := retrieveSwarmInfoNodesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nodesAdded

	err, remoteManagersAdded := retrieveSwarmInfoRemoteManagersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || remoteManagersAdded

	return nil, retAdded
}

func retrieveSwarmInfoClusterFlags(depth int, m *models.SwarmInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	clusterFlagName := fmt.Sprintf("%v.Cluster", cmdPrefix)
	if cmd.Flags().Changed(clusterFlagName) {

		clusterFlagValue := &models.SwarmInfo{}
		err, added := retrieveModelSwarmInfoFlags(depth+1, clusterFlagValue, clusterFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}

func retrieveSwarmInfoControlAvailableFlags(depth int, m *models.SwarmInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	controlAvailableFlagName := fmt.Sprintf("%v.ControlAvailable", cmdPrefix)
	if cmd.Flags().Changed(controlAvailableFlagName) {

		var controlAvailableFlagName string
		if cmdPrefix == "" {
			controlAvailableFlagName = "ControlAvailable"
		} else {
			controlAvailableFlagName = fmt.Sprintf("%v.ControlAvailable", cmdPrefix)
		}

		controlAvailableFlagValue, err := cmd.Flags().GetBool(controlAvailableFlagName)
		if err != nil {
			return err, false
		}
		m.ControlAvailable = &controlAvailableFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveSwarmInfoErrorFlags(depth int, m *models.SwarmInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	errorFlagName := fmt.Sprintf("%v.Error", cmdPrefix)
	if cmd.Flags().Changed(errorFlagName) {

		var errorFlagName string
		if cmdPrefix == "" {
			errorFlagName = "Error"
		} else {
			errorFlagName = fmt.Sprintf("%v.Error", cmdPrefix)
		}

		errorFlagValue, err := cmd.Flags().GetString(errorFlagName)
		if err != nil {
			return err, false
		}
		m.Error = errorFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveSwarmInfoLocalNodeStateFlags(depth int, m *models.SwarmInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	localNodeStateFlagName := fmt.Sprintf("%v.LocalNodeState", cmdPrefix)
	if cmd.Flags().Changed(localNodeStateFlagName) {

		// warning: primitive LocalNodeState LocalNodeState is not supported by go-swagger cli yet

		retAdded = true
	}
	return nil, retAdded
}

func retrieveSwarmInfoManagersFlags(depth int, m *models.SwarmInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	managersFlagName := fmt.Sprintf("%v.Managers", cmdPrefix)
	if cmd.Flags().Changed(managersFlagName) {

		var managersFlagName string
		if cmdPrefix == "" {
			managersFlagName = "Managers"
		} else {
			managersFlagName = fmt.Sprintf("%v.Managers", cmdPrefix)
		}

		managersFlagValue, err := cmd.Flags().GetInt64(managersFlagName)
		if err != nil {
			return err, false
		}
		m.Managers = &managersFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveSwarmInfoNodeAddrFlags(depth int, m *models.SwarmInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	nodeAddrFlagName := fmt.Sprintf("%v.NodeAddr", cmdPrefix)
	if cmd.Flags().Changed(nodeAddrFlagName) {

		var nodeAddrFlagName string
		if cmdPrefix == "" {
			nodeAddrFlagName = "NodeAddr"
		} else {
			nodeAddrFlagName = fmt.Sprintf("%v.NodeAddr", cmdPrefix)
		}

		nodeAddrFlagValue, err := cmd.Flags().GetString(nodeAddrFlagName)
		if err != nil {
			return err, false
		}
		m.NodeAddr = nodeAddrFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveSwarmInfoNodeIDFlags(depth int, m *models.SwarmInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	nodeIdFlagName := fmt.Sprintf("%v.NodeID", cmdPrefix)
	if cmd.Flags().Changed(nodeIdFlagName) {

		var nodeIdFlagName string
		if cmdPrefix == "" {
			nodeIdFlagName = "NodeID"
		} else {
			nodeIdFlagName = fmt.Sprintf("%v.NodeID", cmdPrefix)
		}

		nodeIdFlagValue, err := cmd.Flags().GetString(nodeIdFlagName)
		if err != nil {
			return err, false
		}
		m.NodeID = nodeIdFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveSwarmInfoNodesFlags(depth int, m *models.SwarmInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	nodesFlagName := fmt.Sprintf("%v.Nodes", cmdPrefix)
	if cmd.Flags().Changed(nodesFlagName) {

		var nodesFlagName string
		if cmdPrefix == "" {
			nodesFlagName = "Nodes"
		} else {
			nodesFlagName = fmt.Sprintf("%v.Nodes", cmdPrefix)
		}

		nodesFlagValue, err := cmd.Flags().GetInt64(nodesFlagName)
		if err != nil {
			return err, false
		}
		m.Nodes = &nodesFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveSwarmInfoRemoteManagersFlags(depth int, m *models.SwarmInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	remoteManagersFlagName := fmt.Sprintf("%v.RemoteManagers", cmdPrefix)
	if cmd.Flags().Changed(remoteManagersFlagName) {
		// warning: RemoteManagers array type []*PeerNode is not supported by go-swagger cli yet
	}
	return nil, retAdded
}
