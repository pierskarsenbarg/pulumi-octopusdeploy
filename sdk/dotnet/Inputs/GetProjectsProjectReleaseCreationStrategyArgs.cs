// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetProjectsProjectReleaseCreationStrategyInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("channelId")]
        public Input<string>? ChannelId { get; set; }

        [Input("releaseCreationPackage", required: true)]
        public Input<Inputs.GetProjectsProjectReleaseCreationStrategyReleaseCreationPackageInputArgs> ReleaseCreationPackage { get; set; } = null!;

        [Input("releaseCreationPackageStepId")]
        public Input<string>? ReleaseCreationPackageStepId { get; set; }

        public GetProjectsProjectReleaseCreationStrategyInputArgs()
        {
        }
        public static new GetProjectsProjectReleaseCreationStrategyInputArgs Empty => new GetProjectsProjectReleaseCreationStrategyInputArgs();
    }
}
