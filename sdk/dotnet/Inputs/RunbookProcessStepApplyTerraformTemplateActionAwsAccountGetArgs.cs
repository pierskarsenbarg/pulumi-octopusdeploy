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

    public sealed class RunbookProcessStepApplyTerraformTemplateActionAwsAccountGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("role")]
        public Input<Inputs.RunbookProcessStepApplyTerraformTemplateActionAwsAccountRoleGetArgs>? Role { get; set; }

        [Input("useInstanceRole")]
        public Input<bool>? UseInstanceRole { get; set; }

        [Input("variable")]
        public Input<string>? Variable { get; set; }

        public RunbookProcessStepApplyTerraformTemplateActionAwsAccountGetArgs()
        {
        }
        public static new RunbookProcessStepApplyTerraformTemplateActionAwsAccountGetArgs Empty => new RunbookProcessStepApplyTerraformTemplateActionAwsAccountGetArgs();
    }
}
