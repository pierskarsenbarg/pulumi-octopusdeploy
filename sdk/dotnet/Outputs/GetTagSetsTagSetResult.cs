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
    public sealed class GetTagSetsTagSetResult
    {
        /// <summary>
        /// The description of this tag set.
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
        /// The sort order associated with this resource.
        /// </summary>
        public readonly int SortOrder;
        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        public readonly string SpaceId;

        [OutputConstructor]
        private GetTagSetsTagSetResult(
            string description,

            string id,

            string name,

            int sortOrder,

            string spaceId)
        {
            Description = description;
            Id = id;
            Name = name;
            SortOrder = sortOrder;
            SpaceId = spaceId;
        }
    }
}
