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

    public sealed class RunbookProcessStepApplyTerraformTemplateActionAzureAccountGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("variable")]
        public Input<string>? Variable { get; set; }

        public RunbookProcessStepApplyTerraformTemplateActionAzureAccountGetArgs()
        {
        }
        public static new RunbookProcessStepApplyTerraformTemplateActionAzureAccountGetArgs Empty => new RunbookProcessStepApplyTerraformTemplateActionAzureAccountGetArgs();
    }
}
