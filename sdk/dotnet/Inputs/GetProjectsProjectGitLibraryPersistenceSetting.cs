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

    public sealed class GetProjectsProjectGitLibraryPersistenceSettingArgs : global::Pulumi.InvokeArgs
    {
        [Input("basePath")]
        public string? BasePath { get; set; }

        [Input("defaultBranch")]
        public string? DefaultBranch { get; set; }

        [Input("gitCredentialId", required: true)]
        public string GitCredentialId { get; set; } = null!;

        [Input("protectedBranches")]
        private List<string>? _protectedBranches;
        public List<string> ProtectedBranches
        {
            get => _protectedBranches ?? (_protectedBranches = new List<string>());
            set => _protectedBranches = value;
        }

        [Input("url", required: true)]
        public string Url { get; set; } = null!;

        public GetProjectsProjectGitLibraryPersistenceSettingArgs()
        {
        }
        public static new GetProjectsProjectGitLibraryPersistenceSettingArgs Empty => new GetProjectsProjectGitLibraryPersistenceSettingArgs();
    }
}
