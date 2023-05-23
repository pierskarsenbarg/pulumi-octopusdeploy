// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy
{
    public static class GetProjectGroups
    {
        /// <summary>
        /// Provides information about existing project groups.
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
        ///     var example = Octopusdeploy.GetProjectGroups.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "ProjectGroups-123",
        ///             "ProjectGroups-321",
        ///         },
        ///         PartialName = "Defau",
        ///         Skip = 5,
        ///         Take = 100,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProjectGroupsResult> InvokeAsync(GetProjectGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectGroupsResult>("octopusdeploy:index/getProjectGroups:getProjectGroups", args ?? new GetProjectGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about existing project groups.
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
        ///     var example = Octopusdeploy.GetProjectGroups.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "ProjectGroups-123",
        ///             "ProjectGroups-321",
        ///         },
        ///         PartialName = "Defau",
        ///         Skip = 5,
        ///         Take = 100,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProjectGroupsResult> Invoke(GetProjectGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectGroupsResult>("octopusdeploy:index/getProjectGroups:getProjectGroups", args ?? new GetProjectGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectGroupsArgs : global::Pulumi.InvokeArgs
    {
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
        /// A filter to search by the partial match of a name.
        /// </summary>
        [Input("partialName")]
        public string? PartialName { get; set; }

        [Input("projectGroups")]
        private List<Inputs.GetProjectGroupsProjectGroupArgs>? _projectGroups;

        /// <summary>
        /// A list of project groups that match the filter(s).
        /// </summary>
        public List<Inputs.GetProjectGroupsProjectGroupArgs> ProjectGroups
        {
            get => _projectGroups ?? (_projectGroups = new List<Inputs.GetProjectGroupsProjectGroupArgs>());
            set => _projectGroups = value;
        }

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

        public GetProjectGroupsArgs()
        {
        }
        public static new GetProjectGroupsArgs Empty => new GetProjectGroupsArgs();
    }

    public sealed class GetProjectGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
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
        /// A filter to search by the partial match of a name.
        /// </summary>
        [Input("partialName")]
        public Input<string>? PartialName { get; set; }

        [Input("projectGroups")]
        private InputList<Inputs.GetProjectGroupsProjectGroupInputArgs>? _projectGroups;

        /// <summary>
        /// A list of project groups that match the filter(s).
        /// </summary>
        public InputList<Inputs.GetProjectGroupsProjectGroupInputArgs> ProjectGroups
        {
            get => _projectGroups ?? (_projectGroups = new InputList<Inputs.GetProjectGroupsProjectGroupInputArgs>());
            set => _projectGroups = value;
        }

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

        public GetProjectGroupsInvokeArgs()
        {
        }
        public static new GetProjectGroupsInvokeArgs Empty => new GetProjectGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectGroupsResult
    {
        /// <summary>
        /// An auto-generated identifier that includes the timestamp when this data source was last modified.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        public readonly string? PartialName;
        /// <summary>
        /// A list of project groups that match the filter(s).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectGroupsProjectGroupResult> ProjectGroups;
        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        public readonly int? Skip;
        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        public readonly int? Take;

        [OutputConstructor]
        private GetProjectGroupsResult(
            string id,

            ImmutableArray<string> ids,

            string? partialName,

            ImmutableArray<Outputs.GetProjectGroupsProjectGroupResult> projectGroups,

            int? skip,

            int? take)
        {
            Id = id;
            Ids = ids;
            PartialName = partialName;
            ProjectGroups = projectGroups;
            Skip = skip;
            Take = take;
        }
    }
}
