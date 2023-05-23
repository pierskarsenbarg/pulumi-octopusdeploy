// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Outputs
{

    [OutputType]
    public sealed class LifecyclePhase
    {
        /// <summary>
        /// Environment IDs in this phase that a release is automatically deployed to when it is eligible for this phase
        /// </summary>
        public readonly ImmutableArray<string> AutomaticDeploymentTargets;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// If false a release must be deployed to this phase before it can be deployed to the next phase.
        /// </summary>
        public readonly bool? IsOptionalPhase;
        /// <summary>
        /// The number of units required before a release can enter the next phase. If 0, all environments are required.
        /// </summary>
        public readonly int? MinimumEnvironmentsBeforePromotion;
        /// <summary>
        /// The name of this resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Environment IDs in this phase that a release can be deployed to, but is not automatically deployed to
        /// </summary>
        public readonly ImmutableArray<string> OptionalDeploymentTargets;
        public readonly Outputs.LifecyclePhaseReleaseRetentionPolicy? ReleaseRetentionPolicy;
        public readonly Outputs.LifecyclePhaseTentacleRetentionPolicy? TentacleRetentionPolicy;

        [OutputConstructor]
        private LifecyclePhase(
            ImmutableArray<string> automaticDeploymentTargets,

            string? id,

            bool? isOptionalPhase,

            int? minimumEnvironmentsBeforePromotion,

            string name,

            ImmutableArray<string> optionalDeploymentTargets,

            Outputs.LifecyclePhaseReleaseRetentionPolicy? releaseRetentionPolicy,

            Outputs.LifecyclePhaseTentacleRetentionPolicy? tentacleRetentionPolicy)
        {
            AutomaticDeploymentTargets = automaticDeploymentTargets;
            Id = id;
            IsOptionalPhase = isOptionalPhase;
            MinimumEnvironmentsBeforePromotion = minimumEnvironmentsBeforePromotion;
            Name = name;
            OptionalDeploymentTargets = optionalDeploymentTargets;
            ReleaseRetentionPolicy = releaseRetentionPolicy;
            TentacleRetentionPolicy = tentacleRetentionPolicy;
        }
    }
}