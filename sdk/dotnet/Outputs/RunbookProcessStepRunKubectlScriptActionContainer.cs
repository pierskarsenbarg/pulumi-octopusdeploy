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
    public sealed class RunbookProcessStepRunKubectlScriptActionContainer
    {
        public readonly string? FeedId;
        public readonly string? Image;

        [OutputConstructor]
        private RunbookProcessStepRunKubectlScriptActionContainer(
            string? feedId,

            string? image)
        {
            FeedId = feedId;
            Image = image;
        }
    }
}