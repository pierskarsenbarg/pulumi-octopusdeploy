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

    public sealed class RunbookProcessStepApplyTerraformTemplateActionAwsAccountArgs : global::Pulumi.ResourceArgs
    {
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("role")]
        public Input<Inputs.RunbookProcessStepApplyTerraformTemplateActionAwsAccountRoleArgs>? Role { get; set; }

        [Input("useInstanceRole")]
        public Input<bool>? UseInstanceRole { get; set; }

        [Input("variable")]
        public Input<string>? Variable { get; set; }

        public RunbookProcessStepApplyTerraformTemplateActionAwsAccountArgs()
        {
        }
        public static new RunbookProcessStepApplyTerraformTemplateActionAwsAccountArgs Empty => new RunbookProcessStepApplyTerraformTemplateActionAwsAccountArgs();
    }
}
