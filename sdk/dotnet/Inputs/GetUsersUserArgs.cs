// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetUsersUserInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("canPasswordBeEdited", required: true)]
        public Input<bool> CanPasswordBeEdited { get; set; } = null!;

        /// <summary>
        /// The display name of this resource.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The email address of this resource.
        /// </summary>
        [Input("emailAddress", required: true)]
        public Input<string> EmailAddress { get; set; } = null!;

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("identities", required: true)]
        private InputList<Inputs.GetUsersUserIdentityInputArgs>? _identities;
        public InputList<Inputs.GetUsersUserIdentityInputArgs> Identities
        {
            get => _identities ?? (_identities = new InputList<Inputs.GetUsersUserIdentityInputArgs>());
            set => _identities = value;
        }

        [Input("isActive", required: true)]
        public Input<bool> IsActive { get; set; } = null!;

        [Input("isRequestor", required: true)]
        public Input<bool> IsRequestor { get; set; } = null!;

        [Input("isService", required: true)]
        public Input<bool> IsService { get; set; } = null!;

        /// <summary>
        /// The password associated with this resource.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// The username associated with this resource.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public GetUsersUserInputArgs()
        {
        }
        public static new GetUsersUserInputArgs Empty => new GetUsersUserInputArgs();
    }
}
