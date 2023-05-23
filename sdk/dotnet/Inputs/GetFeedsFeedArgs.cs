// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetFeedsFeedInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessKey", required: true)]
        public Input<string> AccessKey { get; set; } = null!;

        [Input("apiVersion", required: true)]
        public Input<string> ApiVersion { get; set; } = null!;

        [Input("deleteUnreleasedPackagesAfterDays", required: true)]
        public Input<int> DeleteUnreleasedPackagesAfterDays { get; set; } = null!;

        /// <summary>
        /// The number of times a deployment should attempt to download a package from this feed before failing.
        /// </summary>
        [Input("downloadAttempts", required: true)]
        public Input<int> DownloadAttempts { get; set; } = null!;

        /// <summary>
        /// The number of seconds to apply as a linear back off between download attempts.
        /// </summary>
        [Input("downloadRetryBackoffSeconds", required: true)]
        public Input<int> DownloadRetryBackoffSeconds { get; set; } = null!;

        /// <summary>
        /// A filter to search by feed type. Valid feed types are `AwsElasticContainerRegistry`, `BuiltIn`, `Docker`, `GitHub`, `Helm`, `Maven`, `NuGet`, or `OctopusProject`.
        /// </summary>
        [Input("feedType", required: true)]
        public Input<string> FeedType { get; set; } = null!;

        [Input("feedUri", required: true)]
        public Input<string> FeedUri { get; set; } = null!;

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("isEnhancedMode", required: true)]
        public Input<bool> IsEnhancedMode { get; set; } = null!;

        /// <summary>
        /// A short, memorable, unique name for this feed. Example: ACME Builds.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("packageAcquisitionLocationOptions", required: true)]
        private InputList<string>? _packageAcquisitionLocationOptions;
        public InputList<string> PackageAcquisitionLocationOptions
        {
            get => _packageAcquisitionLocationOptions ?? (_packageAcquisitionLocationOptions = new InputList<string>());
            set => _packageAcquisitionLocationOptions = value;
        }

        /// <summary>
        /// The password associated with this resource.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("registryPath", required: true)]
        public Input<string> RegistryPath { get; set; } = null!;

        [Input("secretKey", required: true)]
        public Input<string> SecretKey { get; set; } = null!;

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId", required: true)]
        public Input<string> SpaceId { get; set; } = null!;

        /// <summary>
        /// The username associated with this resource.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public GetFeedsFeedInputArgs()
        {
        }
        public static new GetFeedsFeedInputArgs Empty => new GetFeedsFeedInputArgs();
    }
}
