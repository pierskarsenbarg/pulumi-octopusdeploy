// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Outputs
{

    [OutputType]
    public sealed class ProjectTemplate
    {
        /// <summary>
        /// A default value for the parameter, if applicable. This can be a hard-coded value or a variable reference.
        /// </summary>
        public readonly string? DefaultValue;
        /// <summary>
        /// The display settings for the parameter.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? DisplaySettings;
        /// <summary>
        /// The help presented alongside the parameter input.
        /// </summary>
        public readonly string? HelpText;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The label shown beside the parameter when presented in the deployment process. Example: `Server name`.
        /// </summary>
        public readonly string? Label;
        /// <summary>
        /// The name of the variable set by the parameter. The name can contain letters, digits, dashes and periods. Example: `ServerName`.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private ProjectTemplate(
            string? defaultValue,

            ImmutableDictionary<string, object>? displaySettings,

            string? helpText,

            string? id,

            string? label,

            string name)
        {
            DefaultValue = defaultValue;
            DisplaySettings = displaySettings;
            HelpText = helpText;
            Id = id;
            Label = label;
            Name = name;
        }
    }
}
