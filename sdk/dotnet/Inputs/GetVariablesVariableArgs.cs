// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetVariablesVariableInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this variable.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("encryptedValue", required: true)]
        public Input<string> EncryptedValue { get; set; } = null!;

        /// <summary>
        /// Indicates whether or not this variable is considered editable.
        /// </summary>
        [Input("isEditable", required: true)]
        public Input<bool> IsEditable { get; set; } = null!;

        /// <summary>
        /// Indicates whether or not this resource is considered sensitive and should be kept secret.
        /// </summary>
        [Input("isSensitive", required: true)]
        public Input<bool> IsSensitive { get; set; } = null!;

        [Input("keyFingerprint", required: true)]
        public Input<string> KeyFingerprint { get; set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("ownerId", required: true)]
        public Input<string> OwnerId { get; set; } = null!;

        [Input("pgpKey", required: true)]
        private Input<string>? _pgpKey;
        public Input<string>? PgpKey
        {
            get => _pgpKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _pgpKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("prompts", required: true)]
        private InputList<Inputs.GetVariablesVariablePromptInputArgs>? _prompts;
        public InputList<Inputs.GetVariablesVariablePromptInputArgs> Prompts
        {
            get => _prompts ?? (_prompts = new InputList<Inputs.GetVariablesVariablePromptInputArgs>());
            set => _prompts = value;
        }

        [Input("scopes", required: true)]
        private InputList<Inputs.GetVariablesVariableScopeInputArgs>? _scopes;
        public InputList<Inputs.GetVariablesVariableScopeInputArgs> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<Inputs.GetVariablesVariableScopeInputArgs>());
            set => _scopes = value;
        }

        [Input("sensitiveValue", required: true)]
        private Input<string>? _sensitiveValue;
        public Input<string>? SensitiveValue
        {
            get => _sensitiveValue;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _sensitiveValue = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The type of variable represented by this resource. Valid types are `AmazonWebServicesAccount`, `AzureAccount`, `GoogleCloudAccount`, `Certificate`, `Sensitive`, `String`, or `WorkerPool`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public GetVariablesVariableInputArgs()
        {
        }
        public static new GetVariablesVariableInputArgs Empty => new GetVariablesVariableInputArgs();
    }
}
