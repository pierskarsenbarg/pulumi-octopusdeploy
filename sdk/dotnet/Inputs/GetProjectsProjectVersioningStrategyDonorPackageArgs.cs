// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetProjectsProjectVersioningStrategyDonorPackageInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("deploymentAction")]
        public Input<string>? DeploymentAction { get; set; }

        [Input("packageReference")]
        public Input<string>? PackageReference { get; set; }

        public GetProjectsProjectVersioningStrategyDonorPackageInputArgs()
        {
        }
        public static new GetProjectsProjectVersioningStrategyDonorPackageInputArgs Empty => new GetProjectsProjectVersioningStrategyDonorPackageInputArgs();
    }
}
