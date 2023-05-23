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
    public sealed class DeploymentProcessStepApplyTerraformTemplateActionActionTemplate
    {
        public readonly string? CommunityActionTemplateId;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The version number of this deployment process.
        /// </summary>
        public readonly int? Version;

        [OutputConstructor]
        private DeploymentProcessStepApplyTerraformTemplateActionActionTemplate(
            string? communityActionTemplateId,

            string id,

            int? version)
        {
            CommunityActionTemplateId = communityActionTemplateId;
            Id = id;
            Version = version;
        }
    }
}
