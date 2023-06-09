// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PiersKarsenbarg.Octopusdeploy.Outputs
{

    [OutputType]
    public sealed class RunbookProcessStepDeployPackageActionWindowsService
    {
        public readonly string? Arguments;
        public readonly bool? CreateOrUpdateService;
        public readonly string? CustomAccountName;
        public readonly string? CustomAccountPassword;
        public readonly string? Dependencies;
        public readonly string? Description;
        public readonly string? DisplayName;
        public readonly string ExecutablePath;
        public readonly string? ServiceAccount;
        public readonly string ServiceName;
        public readonly string? StartMode;

        [OutputConstructor]
        private RunbookProcessStepDeployPackageActionWindowsService(
            string? arguments,

            bool? createOrUpdateService,

            string? customAccountName,

            string? customAccountPassword,

            string? dependencies,

            string? description,

            string? displayName,

            string executablePath,

            string? serviceAccount,

            string serviceName,

            string? startMode)
        {
            Arguments = arguments;
            CreateOrUpdateService = createOrUpdateService;
            CustomAccountName = customAccountName;
            CustomAccountPassword = customAccountPassword;
            Dependencies = dependencies;
            Description = description;
            DisplayName = displayName;
            ExecutablePath = executablePath;
            ServiceAccount = serviceAccount;
            ServiceName = serviceName;
            StartMode = startMode;
        }
    }
}
