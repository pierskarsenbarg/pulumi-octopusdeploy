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

    public sealed class LibraryVariableSetTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A default value for the parameter, if applicable. This can be a hard-coded value or a variable reference.
        /// </summary>
        [Input("defaultValue")]
        public Input<string>? DefaultValue { get; set; }

        [Input("displaySettings")]
        private InputMap<object>? _displaySettings;

        /// <summary>
        /// The display settings for the parameter.
        /// </summary>
        public InputMap<object> DisplaySettings
        {
            get => _displaySettings ?? (_displaySettings = new InputMap<object>());
            set => _displaySettings = value;
        }

        /// <summary>
        /// The help presented alongside the parameter input.
        /// </summary>
        [Input("helpText")]
        public Input<string>? HelpText { get; set; }

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The label shown beside the parameter when presented in the deployment process. Example: `Server name`.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// The name of the variable set by the parameter. The name can contain letters, digits, dashes and periods. Example: `ServerName`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public LibraryVariableSetTemplateArgs()
        {
        }
        public static new LibraryVariableSetTemplateArgs Empty => new LibraryVariableSetTemplateArgs();
    }
}
