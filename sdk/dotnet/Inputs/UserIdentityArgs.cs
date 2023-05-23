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

    public sealed class UserIdentityArgs : global::Pulumi.ResourceArgs
    {
        [Input("claims")]
        private InputList<Inputs.UserIdentityClaimArgs>? _claims;
        public InputList<Inputs.UserIdentityClaimArgs> Claims
        {
            get => _claims ?? (_claims = new InputList<Inputs.UserIdentityClaimArgs>());
            set => _claims = value;
        }

        [Input("provider")]
        public Input<string>? Provider { get; set; }

        public UserIdentityArgs()
        {
        }
        public static new UserIdentityArgs Empty => new UserIdentityArgs();
    }
}
