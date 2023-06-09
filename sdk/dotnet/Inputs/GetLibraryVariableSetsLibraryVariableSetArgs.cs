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

    public sealed class GetLibraryVariableSetsLibraryVariableSetInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this library variable set.
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

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId", required: true)]
        public Input<string> SpaceId { get; set; } = null!;

        [Input("templates", required: true)]
        private InputList<Inputs.GetLibraryVariableSetsLibraryVariableSetTemplateInputArgs>? _templates;
        public InputList<Inputs.GetLibraryVariableSetsLibraryVariableSetTemplateInputArgs> Templates
        {
            get => _templates ?? (_templates = new InputList<Inputs.GetLibraryVariableSetsLibraryVariableSetTemplateInputArgs>());
            set => _templates = value;
        }

        [Input("variableSetId", required: true)]
        public Input<string> VariableSetId { get; set; } = null!;

        public GetLibraryVariableSetsLibraryVariableSetInputArgs()
        {
        }
        public static new GetLibraryVariableSetsLibraryVariableSetInputArgs Empty => new GetLibraryVariableSetsLibraryVariableSetInputArgs();
    }
}
