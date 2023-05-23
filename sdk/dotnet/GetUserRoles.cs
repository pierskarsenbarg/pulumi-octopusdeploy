// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy
{
    public static class GetUserRoles
    {
        /// <summary>
        /// Provides information about existing user roles.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Octopusdeploy = Pulumi.Octopusdeploy;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Octopusdeploy.GetUserRoles.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "UserRoles-123",
        ///             "UserRoles-321",
        ///         },
        ///         PartialName = "Administra",
        ///         Skip = 5,
        ///         Take = 100,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUserRolesResult> InvokeAsync(GetUserRolesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserRolesResult>("octopusdeploy:index/getUserRoles:getUserRoles", args ?? new GetUserRolesArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about existing user roles.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Octopusdeploy = Pulumi.Octopusdeploy;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Octopusdeploy.GetUserRoles.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "UserRoles-123",
        ///             "UserRoles-321",
        ///         },
        ///         PartialName = "Administra",
        ///         Skip = 5,
        ///         Take = 100,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetUserRolesResult> Invoke(GetUserRolesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserRolesResult>("octopusdeploy:index/getUserRoles:getUserRoles", args ?? new GetUserRolesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserRolesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        [Input("partialName")]
        public string? PartialName { get; set; }

        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        [Input("skip")]
        public int? Skip { get; set; }

        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        [Input("take")]
        public int? Take { get; set; }

        [Input("userRoles")]
        private List<Inputs.GetUserRolesUserRoleArgs>? _userRoles;

        /// <summary>
        /// A list of user roles that match the filter(s).
        /// </summary>
        public List<Inputs.GetUserRolesUserRoleArgs> UserRoles
        {
            get => _userRoles ?? (_userRoles = new List<Inputs.GetUserRolesUserRoleArgs>());
            set => _userRoles = value;
        }

        public GetUserRolesArgs()
        {
        }
        public static new GetUserRolesArgs Empty => new GetUserRolesArgs();
    }

    public sealed class GetUserRolesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        [Input("partialName")]
        public Input<string>? PartialName { get; set; }

        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        [Input("skip")]
        public Input<int>? Skip { get; set; }

        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        [Input("take")]
        public Input<int>? Take { get; set; }

        [Input("userRoles")]
        private InputList<Inputs.GetUserRolesUserRoleInputArgs>? _userRoles;

        /// <summary>
        /// A list of user roles that match the filter(s).
        /// </summary>
        public InputList<Inputs.GetUserRolesUserRoleInputArgs> UserRoles
        {
            get => _userRoles ?? (_userRoles = new InputList<Inputs.GetUserRolesUserRoleInputArgs>());
            set => _userRoles = value;
        }

        public GetUserRolesInvokeArgs()
        {
        }
        public static new GetUserRolesInvokeArgs Empty => new GetUserRolesInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserRolesResult
    {
        /// <summary>
        /// An auto-generated identifier that includes the timestamp when this data source was last modified.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        public readonly string? PartialName;
        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        public readonly int? Skip;
        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        public readonly int? Take;
        /// <summary>
        /// A list of user roles that match the filter(s).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUserRolesUserRoleResult> UserRoles;

        [OutputConstructor]
        private GetUserRolesResult(
            string id,

            ImmutableArray<string> ids,

            string? partialName,

            int? skip,

            int? take,

            ImmutableArray<Outputs.GetUserRolesUserRoleResult> userRoles)
        {
            Id = id;
            Ids = ids;
            PartialName = partialName;
            Skip = skip;
            Take = take;
            UserRoles = userRoles;
        }
    }
}
