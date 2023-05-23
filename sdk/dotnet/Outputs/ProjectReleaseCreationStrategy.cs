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
    public sealed class ProjectReleaseCreationStrategy
    {
        public readonly string? ChannelId;
        public readonly Outputs.ProjectReleaseCreationStrategyReleaseCreationPackage? ReleaseCreationPackage;
        public readonly string? ReleaseCreationPackageStepId;

        [OutputConstructor]
        private ProjectReleaseCreationStrategy(
            string? channelId,

            Outputs.ProjectReleaseCreationStrategyReleaseCreationPackage? releaseCreationPackage,

            string? releaseCreationPackageStepId)
        {
            ChannelId = channelId;
            ReleaseCreationPackage = releaseCreationPackage;
            ReleaseCreationPackageStepId = releaseCreationPackageStepId;
        }
    }
}