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

    public sealed class GetDeploymentTargetsDeploymentTargetEndpointDestinationInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationType")]
        public Input<string>? DestinationType { get; set; }

        [Input("dropFolderPath")]
        public Input<string>? DropFolderPath { get; set; }

        public GetDeploymentTargetsDeploymentTargetEndpointDestinationInputArgs()
        {
        }
        public static new GetDeploymentTargetsDeploymentTargetEndpointDestinationInputArgs Empty => new GetDeploymentTargetsDeploymentTargetEndpointDestinationInputArgs();
    }
}
