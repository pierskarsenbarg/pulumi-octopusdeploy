// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetVariablesVariablePromptDisplaySettingsInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("controlType", required: true)]
        public Input<string> ControlType { get; set; } = null!;

        [Input("selectOptions")]
        private InputList<Inputs.GetVariablesVariablePromptDisplaySettingsSelectOptionInputArgs>? _selectOptions;
        public InputList<Inputs.GetVariablesVariablePromptDisplaySettingsSelectOptionInputArgs> SelectOptions
        {
            get => _selectOptions ?? (_selectOptions = new InputList<Inputs.GetVariablesVariablePromptDisplaySettingsSelectOptionInputArgs>());
            set => _selectOptions = value;
        }

        public GetVariablesVariablePromptDisplaySettingsInputArgs()
        {
        }
        public static new GetVariablesVariablePromptDisplaySettingsInputArgs Empty => new GetVariablesVariablePromptDisplaySettingsInputArgs();
    }
}
