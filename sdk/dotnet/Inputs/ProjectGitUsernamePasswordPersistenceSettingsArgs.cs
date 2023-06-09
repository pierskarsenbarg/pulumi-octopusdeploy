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

    public sealed class ProjectGitUsernamePasswordPersistenceSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base path associated with these version control settings.
        /// </summary>
        [Input("basePath")]
        public Input<string>? BasePath { get; set; }

        /// <summary>
        /// The default branch associated with these version control settings.
        /// </summary>
        [Input("defaultBranch")]
        public Input<string>? DefaultBranch { get; set; }

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// The password for the Git credential.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("protectedBranches")]
        private InputList<string>? _protectedBranches;

        /// <summary>
        /// A list of protected branch patterns.
        /// </summary>
        public InputList<string> ProtectedBranches
        {
            get => _protectedBranches ?? (_protectedBranches = new InputList<string>());
            set => _protectedBranches = value;
        }

        /// <summary>
        /// The URL associated with these version control settings.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// The username for the Git credential.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public ProjectGitUsernamePasswordPersistenceSettingsArgs()
        {
        }
        public static new ProjectGitUsernamePasswordPersistenceSettingsArgs Empty => new ProjectGitUsernamePasswordPersistenceSettingsArgs();
    }
}
