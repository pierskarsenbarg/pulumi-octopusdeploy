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

    public sealed class GetProjectsProjectTemplateInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("defaultValue")]
        public Input<string>? DefaultValue { get; set; }

        [Input("displaySettings")]
        private InputMap<object>? _displaySettings;
        public InputMap<object> DisplaySettings
        {
            get => _displaySettings ?? (_displaySettings = new InputMap<object>());
            set => _displaySettings = value;
        }

        [Input("helpText")]
        public Input<string>? HelpText { get; set; }

        /// <summary>
        /// An auto-generated identifier that includes the timestamp when this data source was last modified.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// A filter to search by name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetProjectsProjectTemplateInputArgs()
        {
        }
        public static new GetProjectsProjectTemplateInputArgs Empty => new GetProjectsProjectTemplateInputArgs();
    }
}
