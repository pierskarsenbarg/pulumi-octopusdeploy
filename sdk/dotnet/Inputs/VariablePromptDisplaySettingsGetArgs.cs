// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class VariablePromptDisplaySettingsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("controlType", required: true)]
        public Input<string> ControlType { get; set; } = null!;

        [Input("selectOptions")]
        private InputList<Inputs.VariablePromptDisplaySettingsSelectOptionGetArgs>? _selectOptions;
        public InputList<Inputs.VariablePromptDisplaySettingsSelectOptionGetArgs> SelectOptions
        {
            get => _selectOptions ?? (_selectOptions = new InputList<Inputs.VariablePromptDisplaySettingsSelectOptionGetArgs>());
            set => _selectOptions = value;
        }

        public VariablePromptDisplaySettingsGetArgs()
        {
        }
        public static new VariablePromptDisplaySettingsGetArgs Empty => new VariablePromptDisplaySettingsGetArgs();
    }
}
