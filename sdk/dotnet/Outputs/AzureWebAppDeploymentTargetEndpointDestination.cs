// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PiersKarsenbarg.Octopusdeploy.Outputs
{

    [OutputType]
    public sealed class AzureWebAppDeploymentTargetEndpointDestination
    {
        public readonly string? DestinationType;
        public readonly string? DropFolderPath;

        [OutputConstructor]
        private AzureWebAppDeploymentTargetEndpointDestination(
            string? destinationType,

            string? dropFolderPath)
        {
            DestinationType = destinationType;
            DropFolderPath = dropFolderPath;
        }
    }
}
