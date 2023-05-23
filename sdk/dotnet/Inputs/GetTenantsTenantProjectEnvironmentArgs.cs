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

    public sealed class GetTenantsTenantProjectEnvironmentInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("environments", required: true)]
        private InputList<string>? _environments;
        public InputList<string> Environments
        {
            get => _environments ?? (_environments = new InputList<string>());
            set => _environments = value;
        }

        /// <summary>
        /// A filter to search by a project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetTenantsTenantProjectEnvironmentInputArgs()
        {
        }
        public static new GetTenantsTenantProjectEnvironmentInputArgs Empty => new GetTenantsTenantProjectEnvironmentInputArgs();
    }
}
