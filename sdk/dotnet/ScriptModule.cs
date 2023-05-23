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
    /// This resource manages script modules in Octopus Deploy.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Octopusdeploy = Pulumi.Octopusdeploy;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Octopusdeploy.ScriptModule("example", new()
    ///     {
    ///         Description = "A script module to use.",
    ///         Script = new Octopusdeploy.Inputs.ScriptModuleScriptArgs
    ///         {
    ///             Body = @"function Say-Hello()
    /// {
    ///     Write-Output ""Hello, Octopus!""
    /// }
    /// 
    /// ",
    ///             Syntax = "PowerShell",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import octopusdeploy:index/scriptModule:ScriptModule [options] octopusdeploy_script_module.&lt;name&gt; &lt;script-module-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/scriptModule:ScriptModule")]
    public partial class ScriptModule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of this script module.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The script associated with this script module.
        /// </summary>
        [Output("script")]
        public Output<Outputs.ScriptModuleScript> Script { get; private set; } = null!;

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;

        /// <summary>
        /// The variable set ID for this script module.
        /// </summary>
        [Output("variableSetId")]
        public Output<string> VariableSetId { get; private set; } = null!;


        /// <summary>
        /// Create a ScriptModule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScriptModule(string name, ScriptModuleArgs args, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/scriptModule:ScriptModule", name, args ?? new ScriptModuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScriptModule(string name, Input<string> id, ScriptModuleState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/scriptModule:ScriptModule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ScriptModule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScriptModule Get(string name, Input<string> id, ScriptModuleState? state = null, CustomResourceOptions? options = null)
        {
            return new ScriptModule(name, id, state, options);
        }
    }

    public sealed class ScriptModuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this script module.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The script associated with this script module.
        /// </summary>
        [Input("script", required: true)]
        public Input<Inputs.ScriptModuleScriptArgs> Script { get; set; } = null!;

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        /// <summary>
        /// The variable set ID for this script module.
        /// </summary>
        [Input("variableSetId")]
        public Input<string>? VariableSetId { get; set; }

        public ScriptModuleArgs()
        {
        }
        public static new ScriptModuleArgs Empty => new ScriptModuleArgs();
    }

    public sealed class ScriptModuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this script module.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The script associated with this script module.
        /// </summary>
        [Input("script")]
        public Input<Inputs.ScriptModuleScriptGetArgs>? Script { get; set; }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        /// <summary>
        /// The variable set ID for this script module.
        /// </summary>
        [Input("variableSetId")]
        public Input<string>? VariableSetId { get; set; }

        public ScriptModuleState()
        {
        }
        public static new ScriptModuleState Empty => new ScriptModuleState();
    }
}
