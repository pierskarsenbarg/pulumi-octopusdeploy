// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetUsersUserArgs : global::Pulumi.InvokeArgs
    {
        [Input("canPasswordBeEdited", required: true)]
        public bool CanPasswordBeEdited { get; set; }

        /// <summary>
        /// The display name of this resource.
        /// </summary>
        [Input("displayName", required: true)]
        public string DisplayName { get; set; } = null!;

        /// <summary>
        /// The email address of this resource.
        /// </summary>
        [Input("emailAddress", required: true)]
        public string EmailAddress { get; set; } = null!;

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("identities", required: true)]
        private List<Inputs.GetUsersUserIdentityArgs>? _identities;
        public List<Inputs.GetUsersUserIdentityArgs> Identities
        {
            get => _identities ?? (_identities = new List<Inputs.GetUsersUserIdentityArgs>());
            set => _identities = value;
        }

        [Input("isActive", required: true)]
        public bool IsActive { get; set; }

        [Input("isRequestor", required: true)]
        public bool IsRequestor { get; set; }

        [Input("isService", required: true)]
        public bool IsService { get; set; }

        /// <summary>
        /// The password associated with this resource.
        /// </summary>
        [Input("password", required: true)]
        public string Password { get; set; } = null!;

        /// <summary>
        /// The username associated with this resource.
        /// </summary>
        [Input("username", required: true)]
        public string Username { get; set; } = null!;

        public GetUsersUserArgs()
        {
        }
        public static new GetUsersUserArgs Empty => new GetUsersUserArgs();
    }
}