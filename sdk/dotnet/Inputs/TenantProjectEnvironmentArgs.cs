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

    public sealed class TenantProjectEnvironmentArgs : global::Pulumi.ResourceArgs
    {
        [Input("environments", required: true)]
        private InputList<string>? _environments;

        /// <summary>
        /// A list of environment IDs associated with this tenant through a project.
        /// </summary>
        public InputList<string> Environments
        {
            get => _environments ?? (_environments = new InputList<string>());
            set => _environments = value;
        }

        /// <summary>
        /// The project ID associated with this tenant.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public TenantProjectEnvironmentArgs()
        {
        }
        public static new TenantProjectEnvironmentArgs Empty => new TenantProjectEnvironmentArgs();
    }
}
