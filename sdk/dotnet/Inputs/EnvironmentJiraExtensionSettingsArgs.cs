// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class EnvironmentJiraExtensionSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Jira environment type of this Octopus deployment environment. Valid values are `"development"`, `"production"`, `"staging"`, `"testing"`, or `"unmapped"`.
        /// </summary>
        [Input("environmentType", required: true)]
        public Input<string> EnvironmentType { get; set; } = null!;

        public EnvironmentJiraExtensionSettingsArgs()
        {
        }
        public static new EnvironmentJiraExtensionSettingsArgs Empty => new EnvironmentJiraExtensionSettingsArgs();
    }
}