// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PiersKarsenbarg.Octopusdeploy.Outputs
{

    [OutputType]
    public sealed class GetScriptModulesScriptModuleResult
    {
        /// <summary>
        /// The description of this script module.
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
        /// The script associated with this script module.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetScriptModulesScriptModuleScriptResult> Scripts;
        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        public readonly string SpaceId;
        /// <summary>
        /// The variable set ID for this script module.
        /// </summary>
        public readonly string VariableSetId;

        [OutputConstructor]
        private GetScriptModulesScriptModuleResult(
            string description,

            string id,

            string name,

            ImmutableArray<Outputs.GetScriptModulesScriptModuleScriptResult> scripts,

            string spaceId,

            string variableSetId)
        {
            Description = description;
            Id = id;
            Name = name;
            Scripts = scripts;
            SpaceId = spaceId;
            VariableSetId = variableSetId;
        }
    }
}
