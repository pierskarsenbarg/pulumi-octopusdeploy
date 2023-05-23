// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PiersKarsenbarg.Octopusdeploy.Inputs
{

    public sealed class DeploymentProcessStepActionActionTemplateGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("communityActionTemplateId")]
        public Input<string>? CommunityActionTemplateId { get; set; }

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The version number of this deployment process.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public DeploymentProcessStepActionActionTemplateGetArgs()
        {
        }
        public static new DeploymentProcessStepActionActionTemplateGetArgs Empty => new DeploymentProcessStepActionActionTemplateGetArgs();
    }
}
