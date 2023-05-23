// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy
{
    public static class GetMachinePolicies
    {
        /// <summary>
        /// Provides information about existing machine policies.
        /// </summary>
        public static Task<GetMachinePoliciesResult> InvokeAsync(GetMachinePoliciesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMachinePoliciesResult>("octopusdeploy:index/getMachinePolicies:getMachinePolicies", args ?? new GetMachinePoliciesArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about existing machine policies.
        /// </summary>
        public static Output<GetMachinePoliciesResult> Invoke(GetMachinePoliciesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMachinePoliciesResult>("octopusdeploy:index/getMachinePolicies:getMachinePolicies", args ?? new GetMachinePoliciesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMachinePoliciesArgs : global::Pulumi.InvokeArgs
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

        [Input("machinePolicies")]
        private List<Inputs.GetMachinePoliciesMachinePolicyArgs>? _machinePolicies;

        /// <summary>
        /// A list of machine policies that match the filter(s).
        /// </summary>
        public List<Inputs.GetMachinePoliciesMachinePolicyArgs> MachinePolicies
        {
            get => _machinePolicies ?? (_machinePolicies = new List<Inputs.GetMachinePoliciesMachinePolicyArgs>());
            set => _machinePolicies = value;
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

        public GetMachinePoliciesArgs()
        {
        }
        public static new GetMachinePoliciesArgs Empty => new GetMachinePoliciesArgs();
    }

    public sealed class GetMachinePoliciesInvokeArgs : global::Pulumi.InvokeArgs
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

        [Input("machinePolicies")]
        private InputList<Inputs.GetMachinePoliciesMachinePolicyInputArgs>? _machinePolicies;

        /// <summary>
        /// A list of machine policies that match the filter(s).
        /// </summary>
        public InputList<Inputs.GetMachinePoliciesMachinePolicyInputArgs> MachinePolicies
        {
            get => _machinePolicies ?? (_machinePolicies = new InputList<Inputs.GetMachinePoliciesMachinePolicyInputArgs>());
            set => _machinePolicies = value;
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

        public GetMachinePoliciesInvokeArgs()
        {
        }
        public static new GetMachinePoliciesInvokeArgs Empty => new GetMachinePoliciesInvokeArgs();
    }


    [OutputType]
    public sealed class GetMachinePoliciesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of machine policies that match the filter(s).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMachinePoliciesMachinePolicyResult> MachinePolicies;
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
        private GetMachinePoliciesResult(
            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetMachinePoliciesMachinePolicyResult> machinePolicies,

            string? partialName,

            int? skip,

            int? take)
        {
            Id = id;
            Ids = ids;
            MachinePolicies = machinePolicies;
            PartialName = partialName;
            Skip = skip;
            Take = take;
        }
    }
}