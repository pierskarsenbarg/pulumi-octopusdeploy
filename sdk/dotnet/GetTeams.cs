// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy
{
    public static class GetTeams
    {
        /// <summary>
        /// Provides information about existing users.
        /// </summary>
        public static Task<GetTeamsResult> InvokeAsync(GetTeamsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTeamsResult>("octopusdeploy:index/getTeams:getTeams", args ?? new GetTeamsArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about existing users.
        /// </summary>
        public static Output<GetTeamsResult> Invoke(GetTeamsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTeamsResult>("octopusdeploy:index/getTeams:getTeams", args ?? new GetTeamsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTeamsArgs : global::Pulumi.InvokeArgs
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
        /// A filter to include system teams.
        /// </summary>
        [Input("includeSystem")]
        public bool? IncludeSystem { get; set; }

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

        [Input("spaces")]
        private List<string>? _spaces;

        /// <summary>
        /// A filter to search by a list of space IDs.
        /// </summary>
        public List<string> Spaces
        {
            get => _spaces ?? (_spaces = new List<string>());
            set => _spaces = value;
        }

        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        [Input("take")]
        public int? Take { get; set; }

        [Input("teams")]
        private List<Inputs.GetTeamsTeamArgs>? _teams;

        /// <summary>
        /// A list of teams that match the filter(s).
        /// </summary>
        public List<Inputs.GetTeamsTeamArgs> Teams
        {
            get => _teams ?? (_teams = new List<Inputs.GetTeamsTeamArgs>());
            set => _teams = value;
        }

        public GetTeamsArgs()
        {
        }
        public static new GetTeamsArgs Empty => new GetTeamsArgs();
    }

    public sealed class GetTeamsInvokeArgs : global::Pulumi.InvokeArgs
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
        /// A filter to include system teams.
        /// </summary>
        [Input("includeSystem")]
        public Input<bool>? IncludeSystem { get; set; }

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

        [Input("spaces")]
        private InputList<string>? _spaces;

        /// <summary>
        /// A filter to search by a list of space IDs.
        /// </summary>
        public InputList<string> Spaces
        {
            get => _spaces ?? (_spaces = new InputList<string>());
            set => _spaces = value;
        }

        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        [Input("take")]
        public Input<int>? Take { get; set; }

        [Input("teams")]
        private InputList<Inputs.GetTeamsTeamInputArgs>? _teams;

        /// <summary>
        /// A list of teams that match the filter(s).
        /// </summary>
        public InputList<Inputs.GetTeamsTeamInputArgs> Teams
        {
            get => _teams ?? (_teams = new InputList<Inputs.GetTeamsTeamInputArgs>());
            set => _teams = value;
        }

        public GetTeamsInvokeArgs()
        {
        }
        public static new GetTeamsInvokeArgs Empty => new GetTeamsInvokeArgs();
    }


    [OutputType]
    public sealed class GetTeamsResult
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
        /// A filter to include system teams.
        /// </summary>
        public readonly bool? IncludeSystem;
        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        public readonly string? PartialName;
        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        public readonly int? Skip;
        /// <summary>
        /// A filter to search by a list of space IDs.
        /// </summary>
        public readonly ImmutableArray<string> Spaces;
        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        public readonly int? Take;
        /// <summary>
        /// A list of teams that match the filter(s).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTeamsTeamResult> Teams;

        [OutputConstructor]
        private GetTeamsResult(
            string id,

            ImmutableArray<string> ids,

            bool? includeSystem,

            string? partialName,

            int? skip,

            ImmutableArray<string> spaces,

            int? take,

            ImmutableArray<Outputs.GetTeamsTeamResult> teams)
        {
            Id = id;
            Ids = ids;
            IncludeSystem = includeSystem;
            PartialName = partialName;
            Skip = skip;
            Spaces = spaces;
            Take = take;
            Teams = teams;
        }
    }
}
