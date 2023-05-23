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
    public sealed class GetLibraryVariableSetsLibraryVariableSetResult
    {
        /// <summary>
        /// The description of this library variable set.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of this resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        public readonly string SpaceId;
        public readonly ImmutableArray<Outputs.GetLibraryVariableSetsLibraryVariableSetTemplateResult> Templates;
        public readonly string VariableSetId;

        [OutputConstructor]
        private GetLibraryVariableSetsLibraryVariableSetResult(
            string description,

            string id,

            string name,

            string spaceId,

            ImmutableArray<Outputs.GetLibraryVariableSetsLibraryVariableSetTemplateResult> templates,

            string variableSetId)
        {
            Description = description;
            Id = id;
            Name = name;
            SpaceId = spaceId;
            Templates = templates;
            VariableSetId = variableSetId;
        }
    }
}
