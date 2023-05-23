// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class LifecyclePhaseGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("automaticDeploymentTargets")]
        private InputList<string>? _automaticDeploymentTargets;

        /// <summary>
        /// Environment IDs in this phase that a release is automatically deployed to when it is eligible for this phase
        /// </summary>
        public InputList<string> AutomaticDeploymentTargets
        {
            get => _automaticDeploymentTargets ?? (_automaticDeploymentTargets = new InputList<string>());
            set => _automaticDeploymentTargets = value;
        }

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// If false a release must be deployed to this phase before it can be deployed to the next phase.
        /// </summary>
        [Input("isOptionalPhase")]
        public Input<bool>? IsOptionalPhase { get; set; }

        /// <summary>
        /// The number of units required before a release can enter the next phase. If 0, all environments are required.
        /// </summary>
        [Input("minimumEnvironmentsBeforePromotion")]
        public Input<int>? MinimumEnvironmentsBeforePromotion { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("optionalDeploymentTargets")]
        private InputList<string>? _optionalDeploymentTargets;

        /// <summary>
        /// Environment IDs in this phase that a release can be deployed to, but is not automatically deployed to
        /// </summary>
        public InputList<string> OptionalDeploymentTargets
        {
            get => _optionalDeploymentTargets ?? (_optionalDeploymentTargets = new InputList<string>());
            set => _optionalDeploymentTargets = value;
        }

        [Input("releaseRetentionPolicy")]
        public Input<Inputs.LifecyclePhaseReleaseRetentionPolicyGetArgs>? ReleaseRetentionPolicy { get; set; }

        [Input("tentacleRetentionPolicy")]
        public Input<Inputs.LifecyclePhaseTentacleRetentionPolicyGetArgs>? TentacleRetentionPolicy { get; set; }

        public LifecyclePhaseGetArgs()
        {
        }
        public static new LifecyclePhaseGetArgs Empty => new LifecyclePhaseGetArgs();
    }
}