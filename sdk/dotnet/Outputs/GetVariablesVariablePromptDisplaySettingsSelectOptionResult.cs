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
    public sealed class GetVariablesVariablePromptDisplaySettingsSelectOptionResult
    {
        public readonly string DisplayName;
        public readonly string Value;

        [OutputConstructor]
        private GetVariablesVariablePromptDisplaySettingsSelectOptionResult(
            string displayName,

            string value)
        {
            DisplayName = displayName;
            Value = value;
        }
    }
}
