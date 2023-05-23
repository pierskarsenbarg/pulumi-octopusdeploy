// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy
{
    /// <summary>
    /// This resource manages tenant common variables in Octopus Deploy.
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/tenantCommonVariable:TenantCommonVariable")]
    public partial class TenantCommonVariable : global::Pulumi.CustomResource
    {
        [Output("libraryVariableSetId")]
        public Output<string> LibraryVariableSetId { get; private set; } = null!;

        [Output("templateId")]
        public Output<string> TemplateId { get; private set; } = null!;

        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        [Output("value")]
        public Output<string?> Value { get; private set; } = null!;


        /// <summary>
        /// Create a TenantCommonVariable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TenantCommonVariable(string name, TenantCommonVariableArgs args, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/tenantCommonVariable:TenantCommonVariable", name, args ?? new TenantCommonVariableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TenantCommonVariable(string name, Input<string> id, TenantCommonVariableState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/tenantCommonVariable:TenantCommonVariable", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TenantCommonVariable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TenantCommonVariable Get(string name, Input<string> id, TenantCommonVariableState? state = null, CustomResourceOptions? options = null)
        {
            return new TenantCommonVariable(name, id, state, options);
        }
    }

    public sealed class TenantCommonVariableArgs : global::Pulumi.ResourceArgs
    {
        [Input("libraryVariableSetId", required: true)]
        public Input<string> LibraryVariableSetId { get; set; } = null!;

        [Input("templateId", required: true)]
        public Input<string> TemplateId { get; set; } = null!;

        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        [Input("value")]
        public Input<string>? Value { get; set; }

        public TenantCommonVariableArgs()
        {
        }
        public static new TenantCommonVariableArgs Empty => new TenantCommonVariableArgs();
    }

    public sealed class TenantCommonVariableState : global::Pulumi.ResourceArgs
    {
        [Input("libraryVariableSetId")]
        public Input<string>? LibraryVariableSetId { get; set; }

        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public TenantCommonVariableState()
        {
        }
        public static new TenantCommonVariableState Empty => new TenantCommonVariableState();
    }
}