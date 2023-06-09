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
    public sealed class ProjectJiraServiceManagementExtensionSettings
    {
        /// <summary>
        /// The connection identifier associated with the extension settings.
        /// </summary>
        public readonly string ConnectionId;
        /// <summary>
        /// Specifies whether or not this extension is enabled for this project.
        /// </summary>
        public readonly bool IsEnabled;
        /// <summary>
        /// The project name associated with this extension.
        /// </summary>
        public readonly string ServiceDeskProjectName;

        [OutputConstructor]
        private ProjectJiraServiceManagementExtensionSettings(
            string connectionId,

            bool isEnabled,

            string serviceDeskProjectName)
        {
            ConnectionId = connectionId;
            IsEnabled = isEnabled;
            ServiceDeskProjectName = serviceDeskProjectName;
        }
    }
}
