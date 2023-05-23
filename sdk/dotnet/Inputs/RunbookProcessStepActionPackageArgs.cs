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

    public sealed class RunbookProcessStepActionPackageArgs : global::Pulumi.ResourceArgs
    {
        [Input("acquisitionLocation")]
        public Input<string>? AcquisitionLocation { get; set; }

        [Input("extractDuringDeployment")]
        public Input<bool>? ExtractDuringDeployment { get; set; }

        [Input("feedId")]
        public Input<string>? FeedId { get; set; }

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("packageId", required: true)]
        public Input<string> PackageId { get; set; } = null!;

        [Input("properties")]
        private InputMap<string>? _properties;
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        public RunbookProcessStepActionPackageArgs()
        {
        }
        public static new RunbookProcessStepActionPackageArgs Empty => new RunbookProcessStepActionPackageArgs();
    }
}
