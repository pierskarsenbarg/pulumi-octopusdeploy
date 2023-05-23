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

    public sealed class GetLifecyclesLifecycleInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this lifecycle.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("phases", required: true)]
        private InputList<Inputs.GetLifecyclesLifecyclePhaseInputArgs>? _phases;
        public InputList<Inputs.GetLifecyclesLifecyclePhaseInputArgs> Phases
        {
            get => _phases ?? (_phases = new InputList<Inputs.GetLifecyclesLifecyclePhaseInputArgs>());
            set => _phases = value;
        }

        [Input("releaseRetentionPolicies", required: true)]
        private InputList<Inputs.GetLifecyclesLifecycleReleaseRetentionPolicyInputArgs>? _releaseRetentionPolicies;
        public InputList<Inputs.GetLifecyclesLifecycleReleaseRetentionPolicyInputArgs> ReleaseRetentionPolicies
        {
            get => _releaseRetentionPolicies ?? (_releaseRetentionPolicies = new InputList<Inputs.GetLifecyclesLifecycleReleaseRetentionPolicyInputArgs>());
            set => _releaseRetentionPolicies = value;
        }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId", required: true)]
        public Input<string> SpaceId { get; set; } = null!;

        [Input("tentacleRetentionPolicies", required: true)]
        private InputList<Inputs.GetLifecyclesLifecycleTentacleRetentionPolicyInputArgs>? _tentacleRetentionPolicies;
        public InputList<Inputs.GetLifecyclesLifecycleTentacleRetentionPolicyInputArgs> TentacleRetentionPolicies
        {
            get => _tentacleRetentionPolicies ?? (_tentacleRetentionPolicies = new InputList<Inputs.GetLifecyclesLifecycleTentacleRetentionPolicyInputArgs>());
            set => _tentacleRetentionPolicies = value;
        }

        public GetLifecyclesLifecycleInputArgs()
        {
        }
        public static new GetLifecyclesLifecycleInputArgs Empty => new GetLifecyclesLifecycleInputArgs();
    }
}
