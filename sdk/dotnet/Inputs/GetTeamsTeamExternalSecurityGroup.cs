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

    public sealed class GetTeamsTeamExternalSecurityGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("displayIdAndName", required: true)]
        public bool DisplayIdAndName { get; set; }

        [Input("displayName", required: true)]
        public string DisplayName { get; set; } = null!;

        /// <summary>
        /// An auto-generated identifier that includes the timestamp when this data source was last modified.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetTeamsTeamExternalSecurityGroupArgs()
        {
        }
        public static new GetTeamsTeamExternalSecurityGroupArgs Empty => new GetTeamsTeamExternalSecurityGroupArgs();
    }
}
