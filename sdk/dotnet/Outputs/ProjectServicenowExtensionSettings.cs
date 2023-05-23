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
    public sealed class ProjectServicenowExtensionSettings
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
        /// Specifies whether or not this extension will automatically transition the state of a deployment for this project.
        /// </summary>
        public readonly bool IsStateAutomaticallyTransitioned;
        /// <summary>
        /// The name of the standard change template associated with this extension.
        /// </summary>
        public readonly string StandardChangeTemplateName;

        [OutputConstructor]
        private ProjectServicenowExtensionSettings(
            string connectionId,

            bool isEnabled,

            bool isStateAutomaticallyTransitioned,

            string standardChangeTemplateName)
        {
            ConnectionId = connectionId;
            IsEnabled = isEnabled;
            IsStateAutomaticallyTransitioned = isStateAutomaticallyTransitioned;
            StandardChangeTemplateName = standardChangeTemplateName;
        }
    }
}
