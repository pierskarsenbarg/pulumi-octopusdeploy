// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PiersKarsenbarg.Octopusdeploy
{
    /// <summary>
    /// This resource manages an AWS Elastic Container Registry in Octopus Deploy.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Octopusdeploy = PiersKarsenbarg.Octopusdeploy;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Octopusdeploy.AwsElasticContainerRegistry("example", new()
    ///     {
    ///         AccessKey = "access-key",
    ///         Region = "us-east-1",
    ///         SecretKey = "secret-key",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import octopusdeploy:index/awsElasticContainerRegistry:AwsElasticContainerRegistry [options] octopusdeploy_aws_elastic_container_registry.&lt;name&gt; &lt;feed-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/awsElasticContainerRegistry:AwsElasticContainerRegistry")]
    public partial class AwsElasticContainerRegistry : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The AWS access key to use when authenticating against Amazon Web Services.
        /// </summary>
        [Output("accessKey")]
        public Output<string> AccessKey { get; private set; } = null!;

        /// <summary>
        /// A short, memorable, unique name for this feed. Example: ACME Builds.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("packageAcquisitionLocationOptions")]
        public Output<ImmutableArray<string>> PackageAcquisitionLocationOptions { get; private set; } = null!;

        /// <summary>
        /// The AWS region where the registry resides.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The AWS secret key to use when authenticating against Amazon Web Services.
        /// </summary>
        [Output("secretKey")]
        public Output<string> SecretKey { get; private set; } = null!;

        /// <summary>
        /// The space ID associated with this feed.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;


        /// <summary>
        /// Create a AwsElasticContainerRegistry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AwsElasticContainerRegistry(string name, AwsElasticContainerRegistryArgs args, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/awsElasticContainerRegistry:AwsElasticContainerRegistry", name, args ?? new AwsElasticContainerRegistryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AwsElasticContainerRegistry(string name, Input<string> id, AwsElasticContainerRegistryState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/awsElasticContainerRegistry:AwsElasticContainerRegistry", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "secretKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AwsElasticContainerRegistry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AwsElasticContainerRegistry Get(string name, Input<string> id, AwsElasticContainerRegistryState? state = null, CustomResourceOptions? options = null)
        {
            return new AwsElasticContainerRegistry(name, id, state, options);
        }
    }

    public sealed class AwsElasticContainerRegistryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS access key to use when authenticating against Amazon Web Services.
        /// </summary>
        [Input("accessKey", required: true)]
        public Input<string> AccessKey { get; set; } = null!;

        /// <summary>
        /// A short, memorable, unique name for this feed. Example: ACME Builds.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("packageAcquisitionLocationOptions")]
        private InputList<string>? _packageAcquisitionLocationOptions;
        public InputList<string> PackageAcquisitionLocationOptions
        {
            get => _packageAcquisitionLocationOptions ?? (_packageAcquisitionLocationOptions = new InputList<string>());
            set => _packageAcquisitionLocationOptions = value;
        }

        /// <summary>
        /// The AWS region where the registry resides.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("secretKey", required: true)]
        private Input<string>? _secretKey;

        /// <summary>
        /// The AWS secret key to use when authenticating against Amazon Web Services.
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The space ID associated with this feed.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        public AwsElasticContainerRegistryArgs()
        {
        }
        public static new AwsElasticContainerRegistryArgs Empty => new AwsElasticContainerRegistryArgs();
    }

    public sealed class AwsElasticContainerRegistryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS access key to use when authenticating against Amazon Web Services.
        /// </summary>
        [Input("accessKey")]
        public Input<string>? AccessKey { get; set; }

        /// <summary>
        /// A short, memorable, unique name for this feed. Example: ACME Builds.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("packageAcquisitionLocationOptions")]
        private InputList<string>? _packageAcquisitionLocationOptions;
        public InputList<string> PackageAcquisitionLocationOptions
        {
            get => _packageAcquisitionLocationOptions ?? (_packageAcquisitionLocationOptions = new InputList<string>());
            set => _packageAcquisitionLocationOptions = value;
        }

        /// <summary>
        /// The AWS region where the registry resides.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("secretKey")]
        private Input<string>? _secretKey;

        /// <summary>
        /// The AWS secret key to use when authenticating against Amazon Web Services.
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The space ID associated with this feed.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        public AwsElasticContainerRegistryState()
        {
        }
        public static new AwsElasticContainerRegistryState Empty => new AwsElasticContainerRegistryState();
    }
}
