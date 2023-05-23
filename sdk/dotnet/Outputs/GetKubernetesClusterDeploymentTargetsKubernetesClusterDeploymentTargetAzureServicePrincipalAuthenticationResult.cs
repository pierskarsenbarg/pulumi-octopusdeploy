// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Outputs
{

    [OutputType]
    public sealed class GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetAzureServicePrincipalAuthenticationResult
    {
        public readonly string AccountId;
        public readonly string ClusterName;
        public readonly string ClusterResourceGroup;

        [OutputConstructor]
        private GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetAzureServicePrincipalAuthenticationResult(
            string accountId,

            string clusterName,

            string clusterResourceGroup)
        {
            AccountId = accountId;
            ClusterName = clusterName;
            ClusterResourceGroup = clusterResourceGroup;
        }
    }
}