// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class ProjectReleaseCreationStrategyArgs : global::Pulumi.ResourceArgs
    {
        [Input("channelId")]
        public Input<string>? ChannelId { get; set; }

        [Input("releaseCreationPackage")]
        public Input<Inputs.ProjectReleaseCreationStrategyReleaseCreationPackageArgs>? ReleaseCreationPackage { get; set; }

        [Input("releaseCreationPackageStepId")]
        public Input<string>? ReleaseCreationPackageStepId { get; set; }

        public ProjectReleaseCreationStrategyArgs()
        {
        }
        public static new ProjectReleaseCreationStrategyArgs Empty => new ProjectReleaseCreationStrategyArgs();
    }
}