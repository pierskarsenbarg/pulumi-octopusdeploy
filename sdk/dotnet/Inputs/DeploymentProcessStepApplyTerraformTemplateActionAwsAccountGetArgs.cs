// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class DeploymentProcessStepApplyTerraformTemplateActionAwsAccountGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("role")]
        public Input<Inputs.DeploymentProcessStepApplyTerraformTemplateActionAwsAccountRoleGetArgs>? Role { get; set; }

        [Input("useInstanceRole")]
        public Input<bool>? UseInstanceRole { get; set; }

        [Input("variable")]
        public Input<string>? Variable { get; set; }

        public DeploymentProcessStepApplyTerraformTemplateActionAwsAccountGetArgs()
        {
        }
        public static new DeploymentProcessStepApplyTerraformTemplateActionAwsAccountGetArgs Empty => new DeploymentProcessStepApplyTerraformTemplateActionAwsAccountGetArgs();
    }
}
