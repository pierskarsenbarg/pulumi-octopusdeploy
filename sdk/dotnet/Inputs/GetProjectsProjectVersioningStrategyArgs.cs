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

    public sealed class GetProjectsProjectVersioningStrategyInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("donorPackage", required: true)]
        public Input<Inputs.GetProjectsProjectVersioningStrategyDonorPackageInputArgs> DonorPackage { get; set; } = null!;

        [Input("donorPackageStepId", required: true)]
        public Input<string> DonorPackageStepId { get; set; } = null!;

        [Input("template", required: true)]
        public Input<string> Template { get; set; } = null!;

        public GetProjectsProjectVersioningStrategyInputArgs()
        {
        }
        public static new GetProjectsProjectVersioningStrategyInputArgs Empty => new GetProjectsProjectVersioningStrategyInputArgs();
    }
}
