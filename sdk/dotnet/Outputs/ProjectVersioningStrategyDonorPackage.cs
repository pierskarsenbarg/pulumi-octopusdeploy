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
    public sealed class ProjectVersioningStrategyDonorPackage
    {
        public readonly string? DeploymentAction;
        public readonly string? PackageReference;

        [OutputConstructor]
        private ProjectVersioningStrategyDonorPackage(
            string? deploymentAction,

            string? packageReference)
        {
            DeploymentAction = deploymentAction;
            PackageReference = packageReference;
        }
    }
}
