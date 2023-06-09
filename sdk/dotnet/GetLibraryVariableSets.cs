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
    public static class GetLibraryVariableSets
    {
        /// <summary>
        /// Provides information about existing library variable sets.
        /// </summary>
        public static Task<GetLibraryVariableSetsResult> InvokeAsync(GetLibraryVariableSetsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLibraryVariableSetsResult>("octopusdeploy:index/getLibraryVariableSets:getLibraryVariableSets", args ?? new GetLibraryVariableSetsArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about existing library variable sets.
        /// </summary>
        public static Output<GetLibraryVariableSetsResult> Invoke(GetLibraryVariableSetsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLibraryVariableSetsResult>("octopusdeploy:index/getLibraryVariableSets:getLibraryVariableSets", args ?? new GetLibraryVariableSetsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLibraryVariableSetsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter to search by content type.
        /// </summary>
        [Input("contentType")]
        public string? ContentType { get; set; }

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

        [Input("libraryVariableSets")]
        private List<Inputs.GetLibraryVariableSetsLibraryVariableSetArgs>? _libraryVariableSets;

        /// <summary>
        /// A list of library variable sets that match the filter(s).
        /// </summary>
        public List<Inputs.GetLibraryVariableSetsLibraryVariableSetArgs> LibraryVariableSets
        {
            get => _libraryVariableSets ?? (_libraryVariableSets = new List<Inputs.GetLibraryVariableSetsLibraryVariableSetArgs>());
            set => _libraryVariableSets = value;
        }

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

        public GetLibraryVariableSetsArgs()
        {
        }
        public static new GetLibraryVariableSetsArgs Empty => new GetLibraryVariableSetsArgs();
    }

    public sealed class GetLibraryVariableSetsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter to search by content type.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

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

        [Input("libraryVariableSets")]
        private InputList<Inputs.GetLibraryVariableSetsLibraryVariableSetInputArgs>? _libraryVariableSets;

        /// <summary>
        /// A list of library variable sets that match the filter(s).
        /// </summary>
        public InputList<Inputs.GetLibraryVariableSetsLibraryVariableSetInputArgs> LibraryVariableSets
        {
            get => _libraryVariableSets ?? (_libraryVariableSets = new InputList<Inputs.GetLibraryVariableSetsLibraryVariableSetInputArgs>());
            set => _libraryVariableSets = value;
        }

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

        public GetLibraryVariableSetsInvokeArgs()
        {
        }
        public static new GetLibraryVariableSetsInvokeArgs Empty => new GetLibraryVariableSetsInvokeArgs();
    }


    [OutputType]
    public sealed class GetLibraryVariableSetsResult
    {
        /// <summary>
        /// A filter to search by content type.
        /// </summary>
        public readonly string? ContentType;
        /// <summary>
        /// An auto-generated identifier that includes the timestamp when this data source was last modified.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of library variable sets that match the filter(s).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLibraryVariableSetsLibraryVariableSetResult> LibraryVariableSets;
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
        private GetLibraryVariableSetsResult(
            string? contentType,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetLibraryVariableSetsLibraryVariableSetResult> libraryVariableSets,

            string? partialName,

            int? skip,

            int? take)
        {
            ContentType = contentType;
            Id = id;
            Ids = ids;
            LibraryVariableSets = libraryVariableSets;
            PartialName = partialName;
            Skip = skip;
            Take = take;
        }
    }
}
