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

    public sealed class GetProjectsProjectTemplateArgs : global::Pulumi.InvokeArgs
    {
        [Input("defaultValue")]
        public string? DefaultValue { get; set; }

        [Input("displaySettings")]
        private Dictionary<string, object>? _displaySettings;
        public Dictionary<string, object> DisplaySettings
        {
            get => _displaySettings ?? (_displaySettings = new Dictionary<string, object>());
            set => _displaySettings = value;
        }

        [Input("helpText")]
        public string? HelpText { get; set; }

        /// <summary>
        /// An auto-generated identifier that includes the timestamp when this data source was last modified.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("label")]
        public string? Label { get; set; }

        /// <summary>
        /// A filter to search by name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetProjectsProjectTemplateArgs()
        {
        }
        public static new GetProjectsProjectTemplateArgs Empty => new GetProjectsProjectTemplateArgs();
    }
}
