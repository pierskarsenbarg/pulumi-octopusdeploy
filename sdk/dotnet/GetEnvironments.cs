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
    public static class GetEnvironments
    {
        /// <summary>
        /// Provides information about existing environments.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Octopusdeploy = Pulumi.Octopusdeploy;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Octopusdeploy.GetEnvironments.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "Environments-123",
        ///             "Environments-321",
        ///         },
        ///         Name = "Production",
        ///         PartialName = "Produc",
        ///         Skip = 5,
        ///         Take = 100,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEnvironmentsResult> InvokeAsync(GetEnvironmentsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEnvironmentsResult>("octopusdeploy:index/getEnvironments:getEnvironments", args ?? new GetEnvironmentsArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about existing environments.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Octopusdeploy = Pulumi.Octopusdeploy;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Octopusdeploy.GetEnvironments.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "Environments-123",
        ///             "Environments-321",
        ///         },
        ///         Name = "Production",
        ///         PartialName = "Produc",
        ///         Skip = 5,
        ///         Take = 100,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetEnvironmentsResult> Invoke(GetEnvironmentsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnvironmentsResult>("octopusdeploy:index/getEnvironments:getEnvironments", args ?? new GetEnvironmentsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnvironmentsArgs : global::Pulumi.InvokeArgs
    {
        [Input("environments")]
        private List<Inputs.GetEnvironmentsEnvironmentArgs>? _environments;

        /// <summary>
        /// A list of environments that match the filter(s).
        /// </summary>
        public List<Inputs.GetEnvironmentsEnvironmentArgs> Environments
        {
            get => _environments ?? (_environments = new List<Inputs.GetEnvironmentsEnvironmentArgs>());
            set => _environments = value;
        }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A filter to search by name.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        [Input("partialName")]
        public string? PartialName { get; set; }

        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        [Input("skip")]
        public int? Skip { get; set; }

        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        [Input("take")]
        public int? Take { get; set; }

        public GetEnvironmentsArgs()
        {
        }
        public static new GetEnvironmentsArgs Empty => new GetEnvironmentsArgs();
    }

    public sealed class GetEnvironmentsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("environments")]
        private InputList<Inputs.GetEnvironmentsEnvironmentInputArgs>? _environments;

        /// <summary>
        /// A list of environments that match the filter(s).
        /// </summary>
        public InputList<Inputs.GetEnvironmentsEnvironmentInputArgs> Environments
        {
            get => _environments ?? (_environments = new InputList<Inputs.GetEnvironmentsEnvironmentInputArgs>());
            set => _environments = value;
        }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A filter to search by name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        [Input("partialName")]
        public Input<string>? PartialName { get; set; }

        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        [Input("skip")]
        public Input<int>? Skip { get; set; }

        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        [Input("take")]
        public Input<int>? Take { get; set; }

        public GetEnvironmentsInvokeArgs()
        {
        }
        public static new GetEnvironmentsInvokeArgs Empty => new GetEnvironmentsInvokeArgs();
    }


    [OutputType]
    public sealed class GetEnvironmentsResult
    {
        /// <summary>
        /// A list of environments that match the filter(s).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetEnvironmentsEnvironmentResult> Environments;
        /// <summary>
        /// An auto-generated identifier that includes the timestamp when this data source was last modified.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A filter to search by name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        public readonly string? PartialName;
        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        public readonly int? Skip;
        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        public readonly int? Take;

        [OutputConstructor]
        private GetEnvironmentsResult(
            ImmutableArray<Outputs.GetEnvironmentsEnvironmentResult> environments,

            string id,

            ImmutableArray<string> ids,

            string? name,

            string? partialName,

            int? skip,

            int? take)
        {
            Environments = environments;
            Id = id;
            Ids = ids;
            Name = name;
            PartialName = partialName;
            Skip = skip;
            Take = take;
        }
    }
}
