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

    public sealed class GetProjectsProjectGitLibraryPersistenceSettingInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("basePath")]
        public Input<string>? BasePath { get; set; }

        [Input("defaultBranch")]
        public Input<string>? DefaultBranch { get; set; }

        [Input("gitCredentialId", required: true)]
        public Input<string> GitCredentialId { get; set; } = null!;

        [Input("protectedBranches")]
        private InputList<string>? _protectedBranches;
        public InputList<string> ProtectedBranches
        {
            get => _protectedBranches ?? (_protectedBranches = new InputList<string>());
            set => _protectedBranches = value;
        }

        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public GetProjectsProjectGitLibraryPersistenceSettingInputArgs()
        {
        }
        public static new GetProjectsProjectGitLibraryPersistenceSettingInputArgs Empty => new GetProjectsProjectGitLibraryPersistenceSettingInputArgs();
    }
}
