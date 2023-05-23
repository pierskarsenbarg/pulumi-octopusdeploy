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
    /// This resource manages users in Octopus Deploy.
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
    ///     var example = new Octopusdeploy.User("example", new()
    ///     {
    ///         DisplayName = "Bob Smith",
    ///         EmailAddress = "bob.smith@example.com",
    ///         Identities = new[]
    ///         {
    ///             new Octopusdeploy.Inputs.UserIdentityArgs
    ///             {
    ///                 Claims = new[]
    ///                 {
    ///                     new Octopusdeploy.Inputs.UserIdentityClaimArgs
    ///                     {
    ///                         IsIdentifyingClaim = true,
    ///                         Name = "email",
    ///                         Value = "bob.smith@example.com",
    ///                     },
    ///                     new Octopusdeploy.Inputs.UserIdentityClaimArgs
    ///                     {
    ///                         IsIdentifyingClaim = false,
    ///                         Name = "dn",
    ///                         Value = "Bob Smith",
    ///                     },
    ///                 },
    ///                 Provider = "Octopus ID",
    ///             },
    ///         },
    ///         IsActive = true,
    ///         IsService = false,
    ///         Password = "###########",
    ///         Username = "[username]",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import octopusdeploy:index/user:User [options] octopusdeploy_user.&lt;name&gt; &lt;user-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/user:User")]
    public partial class User : global::Pulumi.CustomResource
    {
        [Output("canPasswordBeEdited")]
        public Output<bool> CanPasswordBeEdited { get; private set; } = null!;

        /// <summary>
        /// The display name of this resource.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The email address of this resource.
        /// </summary>
        [Output("emailAddress")]
        public Output<string?> EmailAddress { get; private set; } = null!;

        [Output("identities")]
        public Output<ImmutableArray<Outputs.UserIdentity>> Identities { get; private set; } = null!;

        [Output("isActive")]
        public Output<bool?> IsActive { get; private set; } = null!;

        [Output("isRequestor")]
        public Output<bool> IsRequestor { get; private set; } = null!;

        [Output("isService")]
        public Output<bool?> IsService { get; private set; } = null!;

        /// <summary>
        /// The password associated with this resource.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The username associated with this resource.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/user:User", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display name of this resource.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The email address of this resource.
        /// </summary>
        [Input("emailAddress")]
        public Input<string>? EmailAddress { get; set; }

        [Input("identities")]
        private InputList<Inputs.UserIdentityArgs>? _identities;
        public InputList<Inputs.UserIdentityArgs> Identities
        {
            get => _identities ?? (_identities = new InputList<Inputs.UserIdentityArgs>());
            set => _identities = value;
        }

        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        [Input("isService")]
        public Input<bool>? IsService { get; set; }

        /// <summary>
        /// The password associated with this resource.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The username associated with this resource.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public UserArgs()
        {
        }
        public static new UserArgs Empty => new UserArgs();
    }

    public sealed class UserState : global::Pulumi.ResourceArgs
    {
        [Input("canPasswordBeEdited")]
        public Input<bool>? CanPasswordBeEdited { get; set; }

        /// <summary>
        /// The display name of this resource.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The email address of this resource.
        /// </summary>
        [Input("emailAddress")]
        public Input<string>? EmailAddress { get; set; }

        [Input("identities")]
        private InputList<Inputs.UserIdentityGetArgs>? _identities;
        public InputList<Inputs.UserIdentityGetArgs> Identities
        {
            get => _identities ?? (_identities = new InputList<Inputs.UserIdentityGetArgs>());
            set => _identities = value;
        }

        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        [Input("isRequestor")]
        public Input<bool>? IsRequestor { get; set; }

        [Input("isService")]
        public Input<bool>? IsService { get; set; }

        /// <summary>
        /// The password associated with this resource.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The username associated with this resource.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public UserState()
        {
        }
        public static new UserState Empty => new UserState();
    }
}
